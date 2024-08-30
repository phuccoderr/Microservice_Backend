package com.phuc.reviewservice.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.phuc.reviewservice.dtos.PaginationDto;
import com.phuc.reviewservice.exeptions.ParamValidateException;
import com.phuc.reviewservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import java.util.Set;
import java.util.concurrent.TimeUnit;

@Service
@RequiredArgsConstructor
public class ReviewRedisService implements IReviewRedisService{

    private final RedisTemplate<String, Object> redisTemplate;
    private final ObjectMapper redisObjectMapper;
    @Override
    public void clear() {
        String pattern = "all_reviews:*";
        Set<String> keys = redisTemplate.keys(pattern);
        if (keys != null && !keys.isEmpty()) {
            redisTemplate.delete(keys);
        }
    }

    @Override
    public PaginationDto getAllReviews(Integer page, Integer limit, String sort) throws ParamValidateException, JsonProcessingException {
        Utility.checkSortIsAscOrDesc(sort);

        String key = allReviewsKey(page,limit,sort);
        String json = (String) redisTemplate.opsForValue().get(key);
        return json != null ?
                redisObjectMapper.readValue(json, new TypeReference<PaginationDto>() {})
                : null;
    }

    @Override
    public PaginationDto getAllReviewsByProduct(String productId, Integer page, Integer limit) throws ParamValidateException, JsonProcessingException {

        String key = allReviewsKeyByProduct(productId,page,limit);
        String json = (String) redisTemplate.opsForValue().get(key);
        return json != null ?
                redisObjectMapper.readValue(json, new TypeReference<PaginationDto>() {})
                : null;
    }

    @Override
    public void saveAllReviews(PaginationDto paginationDto, Integer page, Integer limit, String sort) throws JsonProcessingException {
        String key = allReviewsKey(page,limit,sort);
        String json = redisObjectMapper.writeValueAsString(paginationDto);
        redisTemplate.opsForValue().set(key,json,10, TimeUnit.MINUTES);
    }

    @Override
    public void saveAllReviewsByProduct(PaginationDto paginationDto,String proId, Integer page, Integer limit) throws JsonProcessingException {
        String key = allReviewsKeyByProduct(proId,page,limit);
        String json = redisObjectMapper.writeValueAsString(paginationDto);
        redisTemplate.opsForValue().set(key,json,10, TimeUnit.MINUTES);
    }

    private String allReviewsKey(Integer page, Integer limit, String sort) {
        return String.format("all_reviews:%d:%d:%s", page,limit,sort);
    }

    private String allReviewsKeyByProduct(String proId, Integer page, Integer limit) {
        return String.format("all_product_ratings:%s:%d:%d", proId,page,limit);
    }
}
