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
import { Pagination } from 'src/users/dto/pagination.dto';
import { UpdateUserDto } from 'src/users/dto/update-user.dto';
import { JwtAuthGuard } from 'src/auth/guards/jwt-auth.guard';
import { ResposneObject } from 'src/response/response-object.dto';

@Controller('api/v1/users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  // @UseGuards(JwtAuthGuard)
  @Post()
  async createUser(@Body() createUserDto: CreateUserDto) {
    const result: string = await this.usersService.createUser(createUserDto);

    const responseObject: ResposneObject = {
      data: {},
      status: HttpStatus.OK,
      message: result,
    };

    return responseObject;
  }

  @UseGuards(JwtAuthGuard)
  @Get()
  async getUsers(@Query() pagination: Pagination) {
    const users = await this.usersService.getUsers(pagination);

    const responseObject: ResposneObject = {
      data: users,
      status: HttpStatus.OK,
      message: 'Get all users successfully',
    };

    return responseObject;
  }

  @UseGuards(JwtAuthGuard)
  @Get(':id')
  async getUser(@Param('id') _id: string) {
    const user = await this.usersService.getUser(_id);

    const responseObject: ResposneObject = {
      data: user,
      status: HttpStatus.OK,
      message: 'Get user successfully',
    };

    return responseObject;
  }

  @UseGuards(JwtAuthGuard)
  @Patch(':id')
  async updateUser(
    @Param('id') _id: string,
    @Body() updateUserDto: UpdateUserDto,
  ) {
    const result: string = await this.usersService.updateUser(
      _id,
      updateUserDto,
    );
    const responseObject: ResposneObject = {
      data: {},
      status: HttpStatus.OK,
      message: result,
    };
    return responseObject;
  }

  @UseGuards(JwtAuthGuard)
  @Delete(':id')
  async deleteUser(@Param('id') _id: string) {
    const result: string = await this.usersService.deleteUser(_id);

    const responseObject: ResposneObject = {
      data: {},
      status: HttpStatus.OK,
      message: result,
    };
    return responseObject;
  }
}
