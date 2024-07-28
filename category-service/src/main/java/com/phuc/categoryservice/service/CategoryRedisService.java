package com.phuc.categoryservice.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.exceptions.ParamValidateException;
import com.phuc.categoryservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class CategoryRedisService implements ICategoryRedisService{

    private final RedisTemplate<String, Object> redisTemplate;
    private final ObjectMapper redisObjectMapper;
    @Override
    public void clear() {

    }

    @Override
    public List<CategoryDto> getAllCategories(Integer page, Integer limit, String sort)
            throws ParamValidateException {

        Utility.checkSortIsAscOrDesc(sort);

        String key = getKeyFrom(page,limit,sort);
        String json = (String) redisTemplate.opsForValue().get(key);

        return null;
    }

    private String getKeyFrom(Integer page, Integer limit, String sort) {
        return String.format("all_categories:%d:%d:%s", page,limit,sort);
    }
}
