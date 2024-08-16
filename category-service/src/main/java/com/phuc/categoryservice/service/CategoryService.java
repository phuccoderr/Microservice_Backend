package com.phuc.categoryservice.service;

import com.phuc.categoryservice.constants.Constants;
import com.phuc.categoryservice.exceptions.*;
import com.phuc.categoryservice.models.Category;
import com.phuc.categoryservice.repository.CategoryRepository;
import com.phuc.categoryservice.request.RequestCreateCategory;
import com.phuc.categoryservice.request.RequestUpdateCategory;
import com.phuc.categoryservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;


@Service
@RequiredArgsConstructor
public class CategoryService implements ICategoryService {

    private final CategoryRepository repository;

    @Override
    public Category saveCategory(RequestCreateCategory reqCreateCategory)
            throws DataNotFoundException, DataErrorException {
        checkNameUnique(reqCreateCategory.getName());

        Category category = new Category();
        category.setName(reqCreateCategory.getName());
        category.setAlias(Utility.unAccent(category.getName()));
        category.setParent(null);
        category.setStatus(reqCreateCategory.getStatus());

        if (!reqCreateCategory.getParentId().isEmpty()) {
            Category parent = repository.findById(reqCreateCategory.getParentId()).orElseThrow(DataNotFoundException::new);
            category.setParent(parent);
        }

        return repository.save(category);
    }

    @Override
    public Page<Category> getAllCategories(Integer page, Integer limit, String sort, String keyword)
            throws ParamValidateException {

        Utility.checkSortIsAscOrDesc(sort);

        Sort sortDir = Sort.by("name");
        sortDir = sort.equals("asc") ? sortDir.ascending() : sortDir.descending();

        Pageable pageable = PageRequest.of(page - 1,limit,sortDir);
        if (!keyword.isEmpty()) {
            return repository.search(keyword, pageable);
        } else  {
            return repository.findAll(pageable);
        }
    }

    @Override
    public Category getCategory(String id) throws DataNotFoundException {
        return repository.findById(id).orElseThrow(DataNotFoundException::new);
    }

    @Override
    public Category updateCategory(String id, RequestUpdateCategory reqUpdateCategory)
            throws DataNotFoundException, DataErrorException {
        Category category = getCategory(id);

        String cateName = reqUpdateCategory.getName();

        if (!category.getName().equals(cateName)) {
            checkNameUnique(reqUpdateCategory.getName());
            category.setName(cateName);
            category.setAlias(Utility.unAccent(cateName));
        }
        category.setStatus(reqUpdateCategory.getStatus());

        String parentId = reqUpdateCategory.getParent();
        updateParentCategory(parentId,category);

        return repository.save(category);
    }

    @Override
    public void deleteCategory(String id) throws DataNotFoundException, DataErrorException {
        Category category = getCategory(id);
        if (!category.getChildren().isEmpty()) {
            throw new DataErrorException(Constants.DB_HAS_CHILDREN);
        }

        repository.deleteById(category.getId());
    }

    @Override
    public List<String> getChildren(String id) throws DataNotFoundException {
        List<String> categories = new ArrayList<>();

        Category category = getCategory(id);

        categories.add(category.getId());
        hierarchical(category,categories);

        return categories;
    }

    private void hierarchical(Category category, List<String> categories) {
        if (!category.getChildren().isEmpty()) {
            category.getChildren().forEach(child -> {
                categories.add(child.getId());
                hierarchical(child, categories);
            });
        }
    }


    private void checkNameUnique(String name) throws DataErrorException {
        Category category = repository.findByName(name);
        if (category != null) {
            throw new DataErrorException(Constants.DB_ALREADY_EXISTS);
        }
    }


    private void updateParentCategory(String parentId, Category category )
            throws DataNotFoundException, DataErrorException {

        if (parentId.isEmpty()) {
            category.setParent(null);
        } else if (parentId.equals(category.getId())) {
            throw new DataErrorException(Constants.DB_DUPLICATED);
        } else {
            Category parent = getCategory(parentId);
            category.setParent(parent);
        }

    }

}
