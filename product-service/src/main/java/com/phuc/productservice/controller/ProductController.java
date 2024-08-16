package com.phuc.productservice.controller;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.phuc.productservice.constants.Constants;
import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.dtos.PaginationDto;
import com.phuc.productservice.dtos.ProductDto;
import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.exceptions.FuncErrorException;
import com.phuc.productservice.exceptions.ParamValidateException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.request.RequestProduct;
import com.phuc.productservice.response.ResponseObject;
import com.phuc.productservice.service.CategoryService;
import com.phuc.productservice.service.ProductRedisService;
import com.phuc.productservice.service.ProductService;
import com.phuc.productservice.util.Utility;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.util.List;

@RestController
@RequestMapping("/api/v1/products")
@RequiredArgsConstructor
public class ProductController {

    private final ProductService productService;
    private final ProductRedisService productRedisService;
    private final CategoryService categoryService;

    @GetMapping()
    public ResponseEntity<ResponseObject> listByPage(
            @RequestParam(value = "page", defaultValue = "1") Integer page,
            @RequestParam(value = "limit", defaultValue = "10") Integer limit,
            @RequestParam(value = "sort", defaultValue = "asc") String sort,
            @RequestParam(value = "keyword", defaultValue = "") String keyword
    ) throws ParamValidateException, JsonProcessingException {

        if (keyword.isEmpty()) {
            PaginationDto paginationDto = productRedisService.getAllCategories(page, limit, sort);
            if (paginationDto != null ) {
                return new ResponseEntity<>(ResponseObject.builder()
                        .status(HttpStatus.OK.value())
                        .message(Constants.GET_ALL_SUCCESS)
                        .data(paginationDto).build(), HttpStatus.OK);
            }
        }
        Page<Product> pages = productService.getAllProducts(page, limit, sort, keyword);

        List<ProductDto> listDtos = Utility.toListDtos(pages.getContent());

        PaginationDto paginationDto = new PaginationDto(pages,listDtos);
        productRedisService.saveAllCategories(paginationDto,page, limit, sort);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.GET_ALL_SUCCESS)
                .data(paginationDto).build(), HttpStatus.OK);
    }

    @GetMapping("/{id}")
    public ResponseEntity<ResponseObject> getProduct(@PathVariable("id") String id) throws DataErrorException {
        Product product = productService.getProduct(id);

        ProductDto dto = Utility.toDto(product);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.GET_SUCCESS)
                .data(dto).build(), HttpStatus.OK);
    }

    @PostMapping
    public ResponseEntity<ResponseObject> createProduct(
            @RequestPart("product") @Valid RequestProduct requestProduct,
            @RequestParam(value = "main_image", required = false)MultipartFile mainFile,
            @RequestParam(value = "extra_images", required = false) List<MultipartFile> extraFile,
            HttpServletRequest request
            ) throws DataErrorException, FuncErrorException {

        productService.checkNameUnique(requestProduct.getName());

        CategoryDto cateResponse = null;
        if (!requestProduct.getCategoryId().isEmpty()) {
            cateResponse = categoryService.getCategoryById(requestProduct.getCategoryId(), request);
        }

        Product product = productService.createProduct(mainFile,extraFile,requestProduct, cateResponse);

        productRedisService.clear();

        ProductDto dto = Utility.toDto(product);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.CREATED.value())
                .message(Constants.CREATE_SUCCESS)
                .data(dto).build(), HttpStatus.CREATED);
    }

    @PatchMapping("/{id}")
    public ResponseEntity<ResponseObject> updateProduct(
            @PathVariable("id") String id,
            @RequestPart("product") @Valid RequestProduct requestProduct,
            @RequestParam(value = "main_image", required = false)MultipartFile mainFile,
            HttpServletRequest request
    ) throws DataErrorException, FuncErrorException{

        Product productInDB = productService.getProduct(id);

        productService.checkNameUnique(productInDB.getName(),requestProduct.getName());

        CategoryDto cateResponse = null;
        if (!requestProduct.getCategoryId().isEmpty()) {
            cateResponse = categoryService.getCategoryById(requestProduct.getCategoryId(), request);
        }

        Product productUpdated = productService.updateProduct(productInDB, requestProduct, cateResponse, mainFile);

        productRedisService.clear();

        ProductDto dto = Utility.toDto(productUpdated);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.UPDATE_SUCCESS)
                .data(dto).build(), HttpStatus.OK);
    }



    @DeleteMapping("/{id}")
    public ResponseEntity<ResponseObject> deleteProduct(@PathVariable("id") String id) throws DataErrorException {

        productService.deleteProductById(id);

        productRedisService.clear();

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.DELETE_SUCCESS)
                .data("").build(), HttpStatus.OK);
    }

    @PatchMapping("/add_files/{id}")
    public ResponseEntity<ResponseObject> addFilesProduct(
            @PathVariable("id") String id,
            @RequestParam(value = "extra_images", required = false) List<MultipartFile> extraFiles
    ) throws DataErrorException, FuncErrorException {
        Product productInDB = productService.getProduct(id);

        Product productUpdated = productService.addFiles(extraFiles, productInDB);

        ProductDto dto = Utility.toDto(productUpdated);

        productRedisService.clear();

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.ADD_FILES_SUCCESS)
                .data(dto).build(), HttpStatus.OK);
    }

    @DeleteMapping("/delete_files/{id}")
    public ResponseEntity<ResponseObject> deleteFiles(
            @PathVariable("id") String id,
            @RequestBody List<String> listFiles
    ) throws DataErrorException {
        Product productInDB = productService.getProduct(id);

        productService.deleteFiles(listFiles,productInDB);

        productRedisService.clear();

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.DELETE_FILES_SUCCESS)
                .data("").build(), HttpStatus.OK);
    }

    @GetMapping("/c/{cateId}")
    public ResponseEntity<ResponseObject> viewCategoryListByPage(
            @PathVariable(name = "cateId") String cateId,
            @RequestParam(value = "page", defaultValue = "1") Integer page,
            @RequestParam(value = "limit", defaultValue = "10") Integer limit,
            @RequestParam(value = "sort", defaultValue = "asc") String sort,
            @RequestParam(value = "sort_field") String sortField,
            @RequestParam(value = "keyword", defaultValue = "") String keyword,
            HttpServletRequest request
    ) throws ParamValidateException {

        List<String> listCategoriesIds = categoryService.getChildrenCateId(cateId, request);
        Page<Product> pages = productService.getAllProductsByCategory(listCategoriesIds, page, limit, sort, sortField, keyword);

        List<ProductDto> listDtos = Utility.toListDtos(pages.getContent());

        PaginationDto paginationDto = new PaginationDto(pages,listDtos);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.GET_ALL_SUCCESS)
                .data(paginationDto).build(), HttpStatus.OK);
    }

}
