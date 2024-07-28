package com.phuc.categoryservice.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.phuc.categoryservice.dtos.PaginationDto;
import com.phuc.categoryservice.exceptions.ParamValidateException;
import com.phuc.categoryservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import java.util.concurrent.TimeUnit;

@Service
@RequiredArgsConstructor
public class CategoryRedisService implements ICategoryRedisService{

    private final RedisTemplate<String, Object> redisTemplate;
    private final ObjectMapper redisObjectMapper;
    @Override
    public void clear() {
        redisTemplate.getConnectionFactory().getConnection().flushAll();
    }

    @Override
    public PaginationDto getAllCategories(Integer page, Integer limit, String sort)
            throws ParamValidateException, JsonProcessingException {
        Utility.checkSortIsAscOrDesc(sort);

        String key = getKeyFrom(page,limit,sort);
        String json = (String) redisTemplate.opsForValue().get(key);
        return json != null ?
                        redisObjectMapper.readValue(json, new TypeReference<PaginationDto>() {})
                        : null;
    }

    @Override
    public void saveAllCategories(PaginationDto paginationDto, Integer page, Integer limit, String sort)
            throws JsonProcessingException {
        String key = getKeyFrom(page,limit,sort);
        String json = redisObjectMapper.writeValueAsString(paginationDto);
        redisTemplate.opsForValue().set(key,json,10, TimeUnit.MINUTES);
    }

    private String getKeyFrom(Integer page, Integer limit, String sort) {
        return String.format("all_categories:%d:%d:%s", page,limit,sort);
    }
}
