package com.phuc.productservice.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.phuc.productservice.dtos.PaginationDto;
import com.phuc.productservice.exceptions.ParamValidateException;
import com.phuc.productservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import java.util.Set;
import java.util.concurrent.TimeUnit;

@Service
@RequiredArgsConstructor
public class ProductRedisService implements IProductRedisService{

    private final RedisTemplate<String, Object> redisTemplate;
    private final ObjectMapper redisObjectMapper;
    @Override
    public void clear() {
        String pattern = "all_products:*";
        Set<String> keys = redisTemplate.keys(pattern);
        if (keys != null && !keys.isEmpty()) {
            redisTemplate.delete(keys);
        }
    }

    @Override
    public PaginationDto getAllCategories(Integer page, Integer limit, String sort) throws ParamValidateException, JsonProcessingException {
        Utility.checkSortIsAscOrDesc(sort);

        String key = getKeyFrom(page,limit,sort);
        String json = (String) redisTemplate.opsForValue().get(key);
        return json != null ?
                redisObjectMapper.readValue(json, new TypeReference<PaginationDto>() {})
                : null;
    }

    @Override
    public void saveAllCategories(PaginationDto paginationDto, Integer page, Integer limit, String sort) throws JsonProcessingException {
        String key = getKeyFrom(page,limit,sort);
        String json = redisObjectMapper.writeValueAsString(paginationDto);
        redisTemplate.opsForValue().set(key,json,10, TimeUnit.MINUTES);
    }

    private String getKeyFrom(Integer page, Integer limit, String sort) {
        return String.format("all_products:%d:%d:%s", page,limit,sort);
    }
}
