package com.phuc.categoryservice.service;

import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.exceptions.ParamValidateException;

import java.util.List;

public interface ICategoryRedisService {

    void clear();
    List<CategoryDto> getAllCategories(Integer page, Integer limit, String sort) throws ParamValidateException;
}
