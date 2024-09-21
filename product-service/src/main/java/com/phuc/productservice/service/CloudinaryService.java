package com.phuc.productservice.service;

import com.cloudinary.Cloudinary;
import com.phuc.productservice.constants.Constants;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.FuncErrorException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.models.ProductImage;
import com.phuc.productservice.repository.ProductImageRepository;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.*;
import java.util.concurrent.CompletableFuture;

@Service
@RequiredArgsConstructor
public class CloudinaryService implements ICloudinaryService {

    private static final Logger LOGGER = LoggerFactory.getLogger(CloudinaryService.class);

    private final Cloudinary cloudinary;
    private final ProductImageRepository productImageRepository;

    public CompletableFuture<Void> uploadImageAsync(MultipartFile file, String fileName, Product product) {
        return CompletableFuture.runAsync(() -> {
            try {
                Map<String,String> result = cloudinary.uploader()
                        .upload(file.getBytes(), Map.of("public_id", Constants.CLOUD_FOLDER + fileName));
                String publicId = result.get("public_id");
                String url = result.get("secure_url");
                CloudinaryDto cloudinaryDto = new CloudinaryDto(publicId, url);
                productImageRepository.save(new ProductImage(cloudinaryDto,product));
            } catch (Exception ex) {
                LOGGER.warn(ex.getMessage());
                throw new RuntimeException(ex);
            }
        });
    }


    public CloudinaryDto uploadImage(MultipartFile file, String fileName) throws FuncErrorException {
        try {
            Map<String,String> result = cloudinary.uploader()
                    .upload(file.getBytes(), Map.of("public_id", Constants.CLOUD_FOLDER + fileName));
            String publicId = result.get("public_id");
            String url = result.get("secure_url");
            return new CloudinaryDto(publicId,url);
        } catch (Exception ex) {
            LOGGER.info(ex.getMessage());
            throw new FuncErrorException(Constants.UPLOAD_FILE_FAIL);
        }
    }



    public void deleteAsyncImage(String publicId) {
        CompletableFuture.runAsync(() -> {
            try {
                cloudinary.uploader().destroy(publicId, null);
            } catch (IOException e) {
                LOGGER.info(e.getMessage());
                throw new RuntimeException(e);
            }
        });
    }
}
