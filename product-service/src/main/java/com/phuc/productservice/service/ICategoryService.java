package com.phuc.productservice.service;

import com.phuc.productservice.dtos.CategoryDto;
import jakarta.servlet.http.HttpServletRequest;

import java.util.List;

public interface ICategoryService {

    CategoryDto getCategoryById(String categoryId, HttpServletRequest request);

    List<String> getChildrenCateId(String id, HttpServletRequest request);
}
