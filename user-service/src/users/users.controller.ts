import {
  Body,
  Controller,
  Delete,
  Get,
  HttpStatus,
  Param,
  Patch,
  Post,
  Query,
  UseGuards,
} from '@nestjs/common';
import { UsersService } from './users.service';
import { CreateUserDto } from 'src/users/dto/create-user.dto';
import { UpdateUserDto } from 'src/users/dto/update-user.dto';
import { JwtAuthGuard } from 'src/auth/guards/jwt-auth.guard';
import { RedisCacheService } from 'src/redis/redis.service';
import { User } from 'src/users/models/user.schema';
import { allUserKey } from 'src/redis/key';
import { RequestPagination } from 'src/users/dto/request-pagination.dto';
import { ResponseObject } from 'src/response/response-object.dto';
import { PaginationDto } from 'src/users/dto/pagination.dto';

@Controller('api/v1/users')
export class UsersController {
  constructor(
    private readonly usersService: UsersService,
    private readonly redisService: RedisCacheService,
  ) {}

  // @UseGuards(JwtAuthGuard)
  @Post()
  async createUser(
    @Body() createUserDto: CreateUserDto,
  ): Promise<ResponseObject> {
    const result: User = await this.usersService.createUser(createUserDto);

    this.redisService.clearAllUserCache();

    return {
      data: result,
      status: HttpStatus.OK,
      message: 'Success create user',
    };
  }

  @UseGuards(JwtAuthGuard)
  @Get()
  async getUsers(
    @Query() pagination: RequestPagination,
  ): Promise<ResponseObject> {
    const { keyword, page, limit, sort } = pagination;

    if (!keyword) {
      const cachedAllUsers = await this.redisService.get(
        allUserKey(page, limit, sort),
      );
      if (cachedAllUsers) {
        return {
          data: cachedAllUsers,
          status: HttpStatus.OK,
          message: 'Get all users successfully',
        };
      }
    }

    const users = await this.usersService.getUsers(pagination);
    const paginationDto = this.buildPaginationDto(pagination, users);
    this.redisService.set(allUserKey(page, limit, sort), paginationDto);

    return {
      data: paginationDto,
      status: HttpStatus.OK,
      message: 'Get all users successfully',
    };
  }

  @UseGuards(JwtAuthGuard)
  @Get(':id')
  async getUser(@Param('id') _id: string): Promise<ResponseObject> {
    const user = await this.usersService.getUser(_id);

    return {
      data: user,
      status: HttpStatus.OK,
      message: 'Get user successfully',
    };
  }

  @UseGuards(JwtAuthGuard)
  @Patch(':id')
  async updateUser(
    @Param('id') _id: string,
    @Body() updateUserDto: UpdateUserDto,
  ): Promise<ResponseObject> {
    const result: User = await this.usersService.updateUser(_id, updateUserDto);

    this.redisService.clearAllUserCache();

    return {
      data: result,
      status: HttpStatus.OK,
      message: 'Update user success',
    };
  }

  @UseGuards(JwtAuthGuard)
  @Delete(':id')
  async deleteUser(@Param('id') _id: string): Promise<ResponseObject> {
    const result: string = await this.usersService.deleteUser(_id);

    this.redisService.clearAllUserCache();

    return {
      data: {},
      status: HttpStatus.OK,
      message: result,
    };
  }

  private buildPaginationDto(
    pagination: RequestPagination,
    users: User[],
  ): PaginationDto {
    const { page, limit } = pagination;

    return {
      total_items: users.length,
      total_pages: Math.ceil(users.length / limit),
      current_page: parseInt(String(page)),
      start_count: (page - 1) * limit + 1,
      end_count: page * limit > users.length ? users.length : page * limit,
      entities: users,
    };
  }
}
