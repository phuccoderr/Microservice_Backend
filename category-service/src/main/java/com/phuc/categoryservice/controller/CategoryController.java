package com.phuc.categoryservice.controller;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.dtos.PaginationDto;
import com.phuc.categoryservice.exceptions.*;
import com.phuc.categoryservice.models.Category;
import com.phuc.categoryservice.request.RequestCreateCategory;
import com.phuc.categoryservice.request.RequestUpdateCategory;
import com.phuc.categoryservice.response.ResponseObject;
import com.phuc.categoryservice.service.CategoryRedisService;
import com.phuc.categoryservice.service.CategoryService;
import com.phuc.categoryservice.util.Utility;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Collections;
import java.util.List;

@RestController
@RequestMapping("/api/v1/categories")
@RequiredArgsConstructor
public class CategoryController {

    private final CategoryService categoryService;
    private final CategoryRedisService categoryRedisService;

    @PostMapping()
    public ResponseEntity<ResponseObject> createProduct(
            @RequestBody @Valid RequestCreateCategory requestCreateCategory)
            throws DataNotFoundException, DataErrorException {

        Category category = categoryService.saveCategory(requestCreateCategory);
        CategoryDto categoryDto = Utility.toDto(category);
        categoryRedisService.clear();
        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Success create category")
                .data(categoryDto).build(), HttpStatus.CREATED);
    }

    @GetMapping()
    public ResponseEntity<ResponseObject> listByPage(
       @RequestParam(value = "page", defaultValue = "1") Integer page,
       @RequestParam(value = "limit", defaultValue = "10") Integer limit,
       @RequestParam(value = "sort", defaultValue = "asc") String sort,
       @RequestParam(value = "keyword", defaultValue = "") String keyword
    ) throws ParamValidateException, JsonProcessingException {

        if (keyword.isEmpty()) {
            PaginationDto paginationDto = categoryRedisService.getAllCategories(page, limit, sort);
            if (paginationDto != null ) {
                return new ResponseEntity<>(ResponseObject.builder()
                        .status(HttpStatus.OK.value())
                        .message("Get categories successfully")
                        .data(paginationDto).build(), HttpStatus.OK);
            }
        }

        Page<Category> pages = categoryService.getAllCategories(page, limit, sort, keyword);

        List<CategoryDto> listDtos = Utility.toListDtos(pages.getContent());

        PaginationDto paginationDto = new PaginationDto(pages,listDtos);
        categoryRedisService.saveAllCategories(paginationDto,page, limit, sort);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Get categories successfully")
                .data(paginationDto).build(), HttpStatus.OK);
    }

    @GetMapping("/{id}")
    public ResponseEntity<ResponseObject> getCategory(@PathVariable("id") String id)
            throws DataNotFoundException {

        Category category = categoryService.getCategory(id);
        CategoryDto categoryDto = Utility.toDto(category);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Get category successfully")
                .data(categoryDto).build(), HttpStatus.OK);
    }

    @PatchMapping("/{id}")
    public ResponseEntity<ResponseObject> updateCategory(
            @PathVariable("id") String id,
            @RequestBody @Valid RequestUpdateCategory requestUpdateCategory)
            throws DataNotFoundException, DataErrorException {

        Category category = categoryService.updateCategory(id, requestUpdateCategory);
        CategoryDto categoryDto = Utility.toDto(category);
        categoryRedisService.clear();

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Success update category")
                .data(categoryDto).build(), HttpStatus.OK);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<ResponseObject> deleteCategory(@PathVariable("id") String id)
            throws DataNotFoundException, DataErrorException {

        categoryService.deleteCategory(id);
        categoryRedisService.clear();

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Delete category successfully")
                .data(Collections.emptyList()).build(), HttpStatus.OK);

    }
}
