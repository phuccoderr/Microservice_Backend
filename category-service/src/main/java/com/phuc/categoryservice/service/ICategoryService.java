package com.phuc.categoryservice.service;

import com.phuc.categoryservice.exceptions.DataAlreadyExistsException;
import com.phuc.categoryservice.exceptions.DataNotFoundException;
import com.phuc.categoryservice.models.Category;
import com.phuc.categoryservice.request.RequestCreateCategory;
import com.phuc.categoryservice.response.ResponseObject;

public interface ICategoryService {

    Category saveCategory(RequestCreateCategory category) throws DataNotFoundException, DataAlreadyExistsException;
}
