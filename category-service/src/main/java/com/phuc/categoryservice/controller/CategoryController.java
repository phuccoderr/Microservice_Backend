package com.phuc.categoryservice.controller;

import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.exceptions.DataAlreadyExistsException;
import com.phuc.categoryservice.exceptions.DataDuplicatedException;
import com.phuc.categoryservice.exceptions.DataNotFoundException;
import com.phuc.categoryservice.models.Category;
import com.phuc.categoryservice.request.RequestCreateCategory;
import com.phuc.categoryservice.request.RequestUpdateCategory;
import com.phuc.categoryservice.response.ResponseObject;
import com.phuc.categoryservice.service.CategoryService;
import com.phuc.categoryservice.util.Utility;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1/category")
public class CategoryController {

    @Autowired private CategoryService service;

    @GetMapping()
    public ResponseEntity<?> helloworld() {
        return ResponseEntity.ok("Hello world");
    }

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
}
