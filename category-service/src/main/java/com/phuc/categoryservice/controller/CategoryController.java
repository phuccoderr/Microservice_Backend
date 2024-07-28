package com.phuc.categoryservice.controller;

import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.dtos.PaginationDto;
import com.phuc.categoryservice.exceptions.*;
import com.phuc.categoryservice.models.Category;
import com.phuc.categoryservice.request.RequestCreateCategory;
import com.phuc.categoryservice.request.RequestUpdateCategory;
import com.phuc.categoryservice.response.ResponseObject;
import com.phuc.categoryservice.service.CategoryService;
import com.phuc.categoryservice.util.Utility;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Collection;
import java.util.Collections;
import java.util.List;

@RestController
@RequestMapping("/api/v1/category")
public class CategoryController {

    @Autowired private CategoryService service;


    @PostMapping()
    public ResponseEntity<ResponseObject> createProduct(
            @RequestBody @Valid RequestCreateCategory requestCreateCategory)
            throws DataAlreadyExistsException, DataNotFoundException {

        Category category = service.saveCategory(requestCreateCategory);
        CategoryDto categoryDto = Utility.toDto(category);

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
    ) throws ParamValidateException {

        Page<Category> pages = service.getAllCategories(page, limit, sort, keyword);

        Integer pageSize = pages.getSize();
        Integer currentPage = pages.getNumber();
        Integer totalPages = pages.getTotalPages();
        Integer totalItems = (int) pages.getTotalElements();
        int startCount = (currentPage * pageSize) + 1;
        int endCount = Math.min((currentPage * pageSize) + pageSize, totalItems);

        List<CategoryDto> listDtos = Utility.toListDtos(pages.getContent());
        PaginationDto data = PaginationDto.builder()
                .currentPage(currentPage + 1)
                .totalPages(totalPages)
                .totalItems(totalItems)
                .startCount(startCount)
                .endCount(endCount)
                .entities(listDtos)
                .build();

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Get categories successfully")
                .data(data).build(), HttpStatus.OK);
    }

    @GetMapping("/{id}")
    public ResponseEntity<ResponseObject> getCategory(@PathVariable("id") String id) throws DataNotFoundException {
        Category category = service.getCategory(id);
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
            throws DataNotFoundException, DataDuplicatedException, DataAlreadyExistsException {
        Category category = service.updateCategory(id, requestUpdateCategory);
        CategoryDto categoryDto = Utility.toDto(category);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Success update category")
                .data(categoryDto).build(), HttpStatus.OK);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<ResponseObject> deleteCategory(@PathVariable("id") String id)
            throws DataNotFoundException, DataHasChildrenException {

        service.deleteCategory(id);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Delete category successfully")
                .data(Collections.emptyList()).build(), HttpStatus.OK);

    }
}
