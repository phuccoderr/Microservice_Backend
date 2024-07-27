package com.phuc.categoryservice.service;

import com.phuc.categoryservice.exceptions.DataAlreadyExistsException;
import com.phuc.categoryservice.exceptions.DataNotFoundException;
import com.phuc.categoryservice.models.Category;
import com.phuc.categoryservice.repository.CategoryRepository;
import com.phuc.categoryservice.request.RequestCreateCategory;
import com.phuc.categoryservice.response.ResponseObject;
import com.phuc.categoryservice.util.Uitlity;
import org.modelmapper.ModelMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.text.Normalizer;
import java.util.Optional;

@Service
public class CategoryService implements ICategoryService {

    @Autowired
    private CategoryRepository repository;

    ModelMapper mapper = new ModelMapper();

    @Override
    public Category saveCategory(RequestCreateCategory categoryDto) throws DataAlreadyExistsException {
        Category category = repository.findByName(categoryDto.getName());
        if (category != null) {
            throw new DataAlreadyExistsException();
        }

        category = new Category();
        category.setName(categoryDto.getName());
        category.setAlias(Uitlity.unAccent(category.getName()));


        if (!categoryDto.getParentId().isEmpty()) {
            Optional<Category> parent = repository.findById(categoryDto.getParentId());
            category.setParent(null);
            if (parent.isPresent()) {
                category.setParent(parent.get());

            }
        }

        return repository.save(category);
    }


}
