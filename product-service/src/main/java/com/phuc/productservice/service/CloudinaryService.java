package com.phuc.productservice.service;

import com.cloudinary.Cloudinary;
import com.phuc.productservice.constants.Constants;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.FuncErrorException;
import lombok.RequiredArgsConstructor;
import org.springframework.scheduling.annotation.Async;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.*;
import java.util.concurrent.CompletableFuture;

@Service
@RequiredArgsConstructor
@EnableAsync
public class CloudinaryService {

    private final Cloudinary cloudinary;


    public CloudinaryDto uploadImage(MultipartFile file, String fileName) throws FuncErrorException {
        try {
            Map result = cloudinary.uploader()
                    .upload(file.getBytes(), Map.of("public_id", Constants.CLOUD_FOLDER + fileName));
            String publicId = (String) result.get("public_id");
            String url = (String) result.get("secure_url");
            return new CloudinaryDto(publicId,url);
        } catch (Exception ex) {
            throw new FuncErrorException(Constants.UPLOAD_FILE_FAIL);
        }
    }


    @Async
    public void deleteAsyncImage(String publicId) {
        CompletableFuture.runAsync(() -> {
            try {
                cloudinary.uploader().destroy(publicId, null);
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        });
    }
}
