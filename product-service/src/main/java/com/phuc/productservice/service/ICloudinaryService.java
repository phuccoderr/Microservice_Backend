package com.phuc.productservice.service;

import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.FuncErrorException;
import com.phuc.productservice.models.Product;
import org.springframework.web.multipart.MultipartFile;

import java.util.concurrent.CompletableFuture;

public interface ICloudinaryService {

    CompletableFuture<Void> uploadImageAsync(MultipartFile file, String fileName, Product product);
    CloudinaryDto uploadImage(MultipartFile file, String fileName) throws FuncErrorException;
    void deleteAsyncImage(String publicId);
}
