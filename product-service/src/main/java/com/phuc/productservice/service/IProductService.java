package com.phuc.productservice.service;

import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.exceptions.ParamValidateException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.request.RequestCreateProduct;
import com.phuc.productservice.request.RequestProduct;
import org.springframework.data.domain.Page;

public interface IProductService {
    Page<Product> getAllProducts(Integer page, Integer limit, String sort, String keyword) throws ParamValidateException;
    Product getProduct(String proId) throws DataErrorException;
    Product createProduct(Product product, RequestProduct requestProduct, CategoryDto categoryDto);
    void checkNameUnique(String name) throws DataErrorException;
}
