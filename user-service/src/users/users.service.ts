import {
  Inject,
  Injectable,
  InternalServerErrorException,
  NotFoundException,
  UnauthorizedException,
  UnprocessableEntityException,
} from '@nestjs/common';
import { CreateUserDto } from 'src/users/dto/create-user.dto';
import { UsersRepository } from 'src/users/users.repository';
import * as bcrypt from 'bcrypt';
import { Pagination } from 'src/users/dto/pagination.dto';
import { UpdateUserDto } from 'src/users/dto/update-user.dto';
import { RedisCacheService } from 'src/redis/redis.service';
import { allUserKey, limitKey, usersKey } from 'src/redis/key';

@Injectable()
export class UsersService {
  constructor(
    private readonly redisService: RedisCacheService,
    private readonly usersRepository: UsersRepository,
  ) {}

  async createUser(createUserDto: CreateUserDto) {
    try {
      await this.usersRepository.create({
        ...createUserDto,
        password: await bcrypt.hash(createUserDto.password, 10),
      });
    } catch (error) {
      throw new UnprocessableEntityException('Email already exits');
    }
    return 'Success create user';
  }

  async getUsers(query: Pagination) {
    const page = query.page || 1;
    const limit = query.limit || 10;
    const sort = query.sort || 'asc';

    const cachedAllUsers = await this.redisService.get(
      allUserKey(page, limit, sort),
    );

    if (cachedAllUsers) {
      return cachedAllUsers;
    }

    const users = await this.usersRepository.listByPage(page, limit, sort);
    await this.redisService.set(allUserKey(page, limit, sort), users, limitKey);
    return users;
  }

  async getUser(_id: string) {
    const cachedUser = await this.redisService.get(usersKey(_id));

    if (cachedUser) {
      return cachedUser;
    }

    try {
      const user = await this.usersRepository.findOne({ _id }, '-password');
      await this.redisService.set(usersKey(_id), user, limitKey);
      return user;
    } catch (error) {
      throw new NotFoundException('User not found');
    }
  }

  async updateUser(_id: string, updateUserDto: UpdateUserDto) {
    try {
      await this.usersRepository.findOneAndUpdate({ _id }, updateUserDto);
      this.redisService.del(usersKey(_id));
      return 'Update success';
    } catch (error) {
      throw new NotFoundException('User not found');
    }
  }

  async deleteUser(_id: string) {
    try {
      await this.usersRepository.findOneAndDelete({ _id });
      this.redisService.del(usersKey(_id));
      return 'Delete success';
    } catch (error) {
      throw new NotFoundException('User not found');
    }
  }

  async verifyUser(email: string) {
    return await this.usersRepository.findOne({ email }, '');
  }
}
