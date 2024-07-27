package com.phuc.categoryservice.controller;

import com.phuc.categoryservice.exceptions.DataAlreadyExistsException;
import com.phuc.categoryservice.models.Category;
import com.phuc.categoryservice.request.RequestCreateCategory;
import com.phuc.categoryservice.response.ResponseObject;
import com.phuc.categoryservice.service.CategoryService;
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
            @RequestBody @Valid RequestCreateCategory requestCreateCategory) throws DataAlreadyExistsException {
        Category category = service.saveCategory(requestCreateCategory);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Success create category")
                .data(category).build(), HttpStatus.CREATED);
    }
}
