package com.phuc.productservice.controller;

import com.cloudinary.Cloudinary;
import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.dtos.ProductDto;
import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.exceptions.FuncErrorException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.request.RequestCreateProduct;
import com.phuc.productservice.response.ResponseObject;
import com.phuc.productservice.service.CategoryService;
import com.phuc.productservice.service.CloudinaryService;
import com.phuc.productservice.service.ProductService;
import com.phuc.productservice.util.FileUploadUtil;
import com.phuc.productservice.util.Utility;
import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.util.ArrayList;
import java.util.Arrays;

@RestController
@RequestMapping("/api/v1/products")
@RequiredArgsConstructor
public class ProductController {

    private final ProductService productService;
    private final CategoryService categoryService;
    private final CloudinaryService cloudinaryService;

    @PostMapping
    public ResponseEntity<ResponseObject> createProduct(
            @RequestPart("create_product") RequestCreateProduct requestCreateProduct,
            @RequestParam(value = "main_image", required = false)MultipartFile mainFile,
            @RequestParam(value = "extra_image", required = false) MultipartFile[] extraFile,
            HttpServletRequest request
            ) throws DataErrorException, FuncErrorException {

        productService.checkNameUnique(requestCreateProduct.getName());

        cloudinaryService.setMainImage(mainFile,requestCreateProduct);
        cloudinaryService.setExtraImage(extraFile, requestCreateProduct);

        CategoryDto cateResponse = categoryService.getCategoryById(requestCreateProduct.getCategoryId(), request);

        Product product = productService.createProduct(requestCreateProduct, cateResponse);

        ProductDto dto = Utility.toDto(product);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("Success create category")
                .data(dto).build(), HttpStatus.CREATED);
    }


}
