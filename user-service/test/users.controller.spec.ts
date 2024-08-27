import {
  INestApplication,
  NotFoundException,
  ValidationPipe,
} from '@nestjs/common';
import { Test, TestingModule } from '@nestjs/testing';
import { ROLE } from '@src/auth/decorators/role.enum';
import { JwtAuthGuard } from '@src/auth/guards/jwt-auth.guard';
import { RolesAuthGuard } from '@src/auth/guards/roles-auth.guard';
import { RedisCacheService } from '@src/redis/redis.service';
import { CreateUserDto } from '@src/users/dto/create-user.dto';
import { User } from '@src/users/models/user.schema';
import { UsersController } from '@src/users/users.controller';
import { UsersService } from '@src/users/users.service';
import { Types } from 'mongoose';
import * as request from 'supertest';

describe('UsersController', () => {
  let app: INestApplication;
  let usersService: UsersService;

  const userId = new Types.ObjectId();
  const mockUser: User = {
    _id: userId,
    name: 'test',
    password: 'test',
    email: 'test',
    status: true,
  };

  const mockUserService = {
    createUser: jest.fn(),
    getUsers: jest.fn(),
    getUser: jest.fn(),
    updateUser: jest.fn(),
    deleteUser: jest.fn(),
  };
  const mockRedisService = {
    clearAllUserCache: jest.fn(),
    get: jest.fn(),
    set: jest.fn(),
  };

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      controllers: [UsersController],
      providers: [
        { provide: UsersService, useValue: mockUserService },
        { provide: RedisCacheService, useValue: mockRedisService },
      ],
    })
      .overrideGuard(JwtAuthGuard)
      .useValue({ canActivate: () => true })
      .overrideGuard(RolesAuthGuard)
      .useValue({ canActivate: () => true })
      .compile();

    app = moduleFixture.createNestApplication();
    app.useGlobalPipes(new ValidationPipe({ whitelist: true }));
    await app.init();

    usersService = moduleFixture.get<UsersService>(UsersService);
  });

  it('POST /users should 201', async () => {
    const createUserDto: CreateUserDto = {
      email: 'phuc@gmail.com',
      name: 'phuc',
      password: '123456Phuc!',
      status: true,
      roles: [ROLE.ADMIN],
    };

    const response = await request(app.getHttpServer())
      .post('/api/v1/users')
      .send(createUserDto);

    expect(response.status).toEqual(201);
  });

  it('POST /users should 400', async () => {
    const invalidCreateUserDto = {
      email: 'testexample.com', // format email
      name: 123, // is string
      password: 'password', // strong password
      status: 'test', // is boolean
      roles: ['TEST'], // enum ADMIN, USER
    };

    const response = await request(app.getHttpServer())
      .post('/api/v1/users')
      .send(invalidCreateUserDto);

    expect(response.body.message).toContain('email must be an email');
    expect(response.body.message).toContain('name must be a string');
    expect(response.body.message).toContain('password is not strong enough');
    expect(response.body.message).toContain('status must be a boolean value');
    expect(response.body.message).toContain(
      'each value in roles must be one of the following values: ADMIN, USER',
    );
    expect(response.status).toEqual(400);
  });

  it('GET /users should 201', async () => {
    const pagination = {
      page: 1,
      limit: 10,
      sort: 'asc',
      key: 'test',
    };

    jest.spyOn(usersService, 'getUsers').mockResolvedValue([]);

    const response = await request(app.getHttpServer())
      .get('/api/v1/users')
      .query(pagination);

    expect(response.status).toEqual(200);
  });

  it('GET /users should 400', async () => {
    const pagination = {
      page: 'test', // page is number, > 0, max = 100
      limit: 10, // limit is number, > 0, max = 100
      sort: 'test', // sort is asc or desc
      keyword: 123, // key is string
    };

    jest.spyOn(usersService, 'getUsers').mockResolvedValue([]);

    const response = await request(app.getHttpServer())
      .get('/api/v1/users')
      .query(pagination);

    expect(response.body.message).toContain('page must be a positive number');
    expect(response.body.message).toContain('page must be an integer number');
    expect(response.body.message).toContain(
      'sort must be one of the following values: asc, desc',
    );
    expect(response.status).toEqual(400);
  });

  it('GET /users/:id should 200', async () => {
    jest.spyOn(usersService, 'getUser').mockResolvedValue(mockUser);

    const response = await request(app.getHttpServer()).get(
      `/api/v1/users/${userId}`,
    );

    expect(response.status).toEqual(200);
  });

  it('GET /users/:id should 404', async () => {
    jest
      .spyOn(usersService, 'getUser')
      .mockRejectedValue(new NotFoundException());

    const response = await request(app.getHttpServer()).get(
      `/api/v1/users/${userId}`,
    );

    expect(response.status).toEqual(404);
  });

  it('PATCH /users/:id should 200', async () => {
    const update = {
      name: 'test',
      roles: [ROLE.ADMIN],
    };

    jest.spyOn(usersService, 'updateUser').mockResolvedValue(mockUser);

    const response = await request(app.getHttpServer())
      .patch(`/api/v1/users/${userId}`)
      .send(update);

    expect(response.status).toEqual(200);
  });

  it('PATCH /users/:id should 400', async () => {
    const update = {
      name: '', // is string, not empty
      roles: ['TEST'], // enum ADMIN, USER
    };

    jest.spyOn(usersService, 'updateUser').mockResolvedValue(mockUser);

    const response = await request(app.getHttpServer())
      .patch(`/api/v1/users/${userId}`)
      .send(update);

    expect(response.body.message).toContain('name should not be empty');
    expect(response.body.message).toContain(
      'each value in roles must be one of the following values: ADMIN, USER',
    );
    expect(response.status).toEqual(400);
  });

  it('PATCH /users/:id should 404', async () => {
    const update = {
      name: 'test',
      roles: [ROLE.ADMIN],
    };

    jest
      .spyOn(usersService, 'updateUser')
      .mockRejectedValue(new NotFoundException());

    const response = await request(app.getHttpServer())
      .patch(`/api/v1/users/${userId}`)
      .send(update);

    expect(response.status).toEqual(404);
  });

  it('DELETE /users/:id should 200', async () => {
    jest.spyOn(usersService, 'deleteUser').mockResolvedValue(undefined);

    const response = await request(app.getHttpServer()).delete(
      `/api/v1/users/${userId}`,
    );

    expect(response.status).toEqual(200);
  });

  it('DELETE /users/:id should 404', async () => {
    jest
      .spyOn(usersService, 'deleteUser')
      .mockRejectedValue(new NotFoundException());

    const response = await request(app.getHttpServer()).delete(
      `/api/v1/users/${userId}`,
    );

    expect(response.status).toEqual(404);
  });
});
