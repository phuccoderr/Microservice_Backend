import {
  Body,
  Controller,
  Delete,
  Get,
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

@Controller('api/v1/users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  // @UseGuards(JwtAuthGuard)
  @Post()
  async createUser(@Body() createUserDto: CreateUserDto) {
    return this.usersService.createUser(createUserDto);
  }

  // @UseGuards(JwtAuthGuard)
  @Get()
  async getUsers(@Query() pagination: Pagination) {
    return this.usersService.getUsers(pagination);
  }

  // @UseGuards(JwtAuthGuard)
  @Get(':id')
  async getUser(@Param('id') _id: string) {
    return this.usersService.getUser(_id);
  }

  // @UseGuards(JwtAuthGuard)
  @Patch(':id')
  async updateUser(
    @Param('id') _id: string,
    @Body() updateUserDto: UpdateUserDto,
  ) {
    return this.usersService.updateUser(_id, updateUserDto);
  }

  // @UseGuards(JwtAuthGuard)
  @Delete(':id')
  async deleteUser(@Param('id') _id: string) {
    return this.usersService.deleteUser(_id);
  }
}
