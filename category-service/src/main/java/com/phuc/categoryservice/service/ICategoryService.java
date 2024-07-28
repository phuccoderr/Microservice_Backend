package com.phuc.categoryservice.service;

import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.exceptions.*;
import com.phuc.categoryservice.models.Category;
import com.phuc.categoryservice.request.RequestCreateCategory;
import com.phuc.categoryservice.request.RequestUpdateCategory;
import com.phuc.categoryservice.response.ResponseObject;
import org.springframework.data.domain.Page;

import java.util.List;

public interface ICategoryService {

    Category saveCategory(RequestCreateCategory category) throws DataNotFoundException, DataAlreadyExistsException;
    Page<Category> getAllCategories(Integer page, Integer limit, String sort, String keyword) throws ParamValidateException;
    Category getCategory(String id) throws DataNotFoundException;
    Category updateCategory(String id, RequestUpdateCategory reqUpdateCategory) throws DataNotFoundException, DataDuplicatedException, DataAlreadyExistsException;
    void deleteCategory(String id) throws DataNotFoundException, DataHasChildrenException;
}
