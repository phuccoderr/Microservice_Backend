import {
  BadRequestException,
  Injectable,
  Logger,
  NotFoundException,
  OnModuleInit,
  UnprocessableEntityException,
} from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UsersRepository } from './repository/users.repository';
import * as bcrypt from 'bcrypt';
import { UpdateUserDto } from './dto/update-user.dto';
import { User } from './models/user.schema';
import { RequestPaginationDto } from './dto/request-pagination.dto';
import { DATABASE_CONST } from '@src/constants/db-constants';
import { ROLE } from '@src/auth/decorators/role.enum';
import { USER_CONSTANTS } from '@src/constants/user-constants';

@Injectable()
export class UsersService implements OnModuleInit {
  private logger = new Logger(UsersService.name);
  constructor(private readonly usersRepository: UsersRepository) {}

  async createUser(createUserDto: CreateUserDto): Promise<User> {
    try {
      return await this.usersRepository.create({
        ...createUserDto,
        password: await bcrypt.hash(createUserDto.password, 10),
      });
    } catch (error) {
      this.logger.error(error);
      throw new UnprocessableEntityException(error.message);
    }
  }

  async getUsers(query: RequestPaginationDto) {
    const { keyword, page, limit, sort } = query;

    if (keyword) {
      const filter = { email: keyword, name: keyword };
      return await this.usersRepository.search(page, limit, sort, filter);
    }

    return await this.usersRepository.listByPage(page, limit, sort);
  }

  async getUser(_id: string) {
    try {
      return await this.usersRepository.findOne({ _id }, '-password');
    } catch (error) {
      this.logger.warn('User does not exist with id ' + _id);
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }
  }

  async updateUser(_id: string, updateUserDto: UpdateUserDto): Promise<User> {
    try {
      return await this.usersRepository.findOneAndUpdate(
        { _id },
        updateUserDto,
      );
    } catch (error) {
      this.logger.warn('User does not exist with id ' + _id);
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }
  }

  async deleteUser(_id: string): Promise<void> {
    const user = await this.usersRepository.findOneAndDelete({ _id });
    if (!user) {
      this.logger.warn('User does not exist with id ' + _id);
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }
  }

  async updateStatus(_id: string, status: string) {

    if(status != 'true' && status != 'false') {
      this.logger.warn('Status is true or false!');
      throw new BadRequestException(USER_CONSTANTS.STATUS);
    }

    try {
      await this.usersRepository.findOneAndUpdate({_id}, {status});
    } catch (error) {
      this.logger.warn('User does not exist with id ' + _id);
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }
  }

  async onModuleInit(): Promise<void> {
    const createUserAdmin: CreateUserDto = {
      email: 'phuc@gmail.com',
      name: 'phuc',
      password: '123456Phuc!',
      status: true,
      roles: [ROLE.ADMIN],
    };
    const initInDB = await this.usersRepository.findOne(
      { email: createUserAdmin.email },
      '',
    );
    if (!initInDB) {
      await this.usersRepository.create({
        ...createUserAdmin,
        password: await bcrypt.hash(createUserAdmin.password, 10),
      });
    }
  }
}
