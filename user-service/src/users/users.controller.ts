import {
  Body,
  Controller,
  Delete,
  Get,
  HttpCode,
  HttpStatus,
  Param,
  Patch,
  Post,
  Query,
  UseGuards,
} from '@nestjs/common';
import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { JwtAuthGuard } from '../auth/guards/jwt-auth.guard';
import { RedisCacheService } from '../redis/redis.service';
import { User } from './models/user.schema';
import { allUserKey } from '../redis/key';
import { RequestPaginationDto } from './dto/request-pagination.dto';
import { ResponseObject } from '../response/response-object.dto';
import { ResponsePaginationDto } from './dto/response-pagination.dto';
import { USER_CONSTANTS } from '@src/constants/user-constants';
import { ROLE } from '@src/auth/decorators/role.enum';
import { Roles } from '@src/auth/decorators/roles.decorator';
import { RolesAuthGuard } from '@src/auth/guards/roles-auth.guard';

@Controller('api/v1/users')
export class UsersController {
  constructor(
    private readonly usersService: UsersService,
    private readonly redisService: RedisCacheService,
  ) {}

  @Post()
  @UseGuards(JwtAuthGuard, RolesAuthGuard)
  @Roles(ROLE.ADMIN)
  async createUser(
    @Body() createUserDto: CreateUserDto,
  ): Promise<ResponseObject> {
    const result: User = await this.usersService.createUser(createUserDto);

    this.redisService.clearAllUserCache();

    return {
      data: result,
      status: HttpStatus.CREATED,
      message: USER_CONSTANTS.CREATE,
    };
  }

  @UseGuards(JwtAuthGuard)
  @Get()
  async getUsers(
    @Query() pagination: RequestPaginationDto,
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
          message: USER_CONSTANTS.GET_ALL,
        };
      }
    }

    const users = await this.usersService.getUsers(pagination);
    const paginationDto = this.buildPaginationDto(pagination, users);
    if (!keyword) {
      this.redisService.set(allUserKey(page, limit, sort), paginationDto);
    }

    return {
      data: paginationDto,
      status: HttpStatus.OK,
      message: USER_CONSTANTS.GET_ALL,
    };
  }

  @UseGuards(JwtAuthGuard)
  @Get(':id')
  async getUser(@Param('id') _id: string): Promise<ResponseObject> {
    const user = await this.usersService.getUser(_id);

    return {
      data: user,
      status: HttpStatus.OK,
      message: USER_CONSTANTS.GET,
    };
  }

  @UseGuards(JwtAuthGuard, RolesAuthGuard)
  @Roles(ROLE.ADMIN)
  @Patch(':id')
  @HttpCode(HttpStatus.OK)
  async updateUser(
    @Param('id') _id: string,
    @Body() updateUserDto: UpdateUserDto,
  ): Promise<ResponseObject> {
    const result: User = await this.usersService.updateUser(_id, updateUserDto);

    this.redisService.clearAllUserCache();

    return {
      data: result,
      status: HttpStatus.OK,
      message: USER_CONSTANTS.UPDATE,
    };
  }

  @UseGuards(JwtAuthGuard, RolesAuthGuard)
  @Roles(ROLE.ADMIN)
  @Delete(':id')
  async deleteUser(@Param('id') _id: string): Promise<ResponseObject> {
    await this.usersService.deleteUser(_id);

    this.redisService.clearAllUserCache();

    return {
      data: {},
      status: HttpStatus.OK,
      message: USER_CONSTANTS.DELETE,
    };
  }

  @UseGuards(JwtAuthGuard, RolesAuthGuard)
  @Roles(ROLE.ADMIN)
  @Patch('/:id/status/:status')
  async updateStatus(
    @Param('id') _id: string,
    @Param('status') status: string,
  ): Promise<ResponseObject> {
    await this.usersService.updateStatus(_id, status);

    this.redisService.clearAllUserCache();

    return {
      data: status,
      status: HttpStatus.OK,
      message: USER_CONSTANTS.UPDATE_STATUS,
    };
  }

  private buildPaginationDto(
    pagination: RequestPaginationDto,
    users: User[],
  ): ResponsePaginationDto {
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
