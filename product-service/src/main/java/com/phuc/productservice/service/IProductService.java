package com.phuc.productservice.service;

import com.phuc.productservice.dtos.CategoryDto;
import com.phuc.productservice.exceptions.DataErrorException;
import com.phuc.productservice.exceptions.DataNotFoundException;
import com.phuc.productservice.exceptions.FuncErrorException;
import com.phuc.productservice.exceptions.ParamValidateException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.request.RequestProduct;
import org.springframework.data.domain.Page;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.List;
import java.util.concurrent.CompletableFuture;

public interface IProductService {
    Page<Product> getAllProducts(Integer page, Integer limit, String sort, String keyword) throws ParamValidateException;
    Product getProduct(String proId) throws DataErrorException, DataNotFoundException;

    Product createProduct(
            MultipartFile mainImage,
            List<MultipartFile> extraImages,
            RequestProduct requestProduct,
            CategoryDto categoryDto
    ) throws FuncErrorException;

    Product updateProduct(
            Product proInDB,
            RequestProduct requestProduct,
            CategoryDto categoryDto,
            MultipartFile mainImage
    ) throws FuncErrorException, IOException;

    void deleteProductById(String proId) throws DataErrorException, DataNotFoundException;
    Page<Product> getAllProductsByCategory(
            List<String> listCategoryIds,
            Integer page,
            Integer limit,
            String sort,
            String sortField,
            String keyword
    ) throws ParamValidateException;
    void setExtraImage(List<MultipartFile> extraFile,Product product, String socketId);
    void deleteFiles(List<String> listFiles, Product product);

    void checkNameUnique(String oldName, String newName) throws DataErrorException;

    void checkNameUnique(String name) throws DataErrorException;
}
