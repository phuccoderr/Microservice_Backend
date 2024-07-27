package com.phuc.categoryservice.service;

import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.exceptions.DataAlreadyExistsException;
import com.phuc.categoryservice.exceptions.DataDuplicatedException;
import com.phuc.categoryservice.exceptions.DataNotFoundException;
import com.phuc.categoryservice.models.Category;
import com.phuc.categoryservice.request.RequestCreateCategory;
import com.phuc.categoryservice.request.RequestUpdateCategory;
import com.phuc.categoryservice.response.ResponseObject;

public interface ICategoryService {

    Category saveCategory(RequestCreateCategory category) throws DataNotFoundException, DataAlreadyExistsException;
    Category getCategory(String id) throws DataNotFoundException;

    Category updateCategory(String id, RequestUpdateCategory reqUpdateCategory) throws DataNotFoundException, DataDuplicatedException, DataAlreadyExistsException;
}
