import { CACHE_MANAGER, CacheStore } from '@nestjs/cache-manager';
import { Inject, Injectable } from '@nestjs/common';
@Injectable()
export class RedisCacheService {
  constructor(@Inject(CACHE_MANAGER) private cacheManager: CacheStore) {}

  async get(key: string) {
    return await this.cacheManager.get(key);
  }

  async set(key: string, value: object, limit: number) {
    await this.cacheManager.set(key, value, { ttl: limit });
  }

  async del(key: string) {
    await this.cacheManager.del(key);
  }
}
