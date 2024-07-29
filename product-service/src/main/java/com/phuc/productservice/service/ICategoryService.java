package com.phuc.productservice.service;

import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.response.ResponseObject;
import jakarta.servlet.http.HttpServletRequest;

public interface ICategoryService {

    CategoryDto getCategoryById(String categoryId, HttpServletRequest request);
}
