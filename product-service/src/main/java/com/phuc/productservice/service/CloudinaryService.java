package com.phuc.productservice.service;

import com.cloudinary.Cloudinary;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.FuncErrorException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.models.ProductImage;
import com.phuc.productservice.request.RequestProduct;
import com.phuc.productservice.util.FileUploadUtil;
import lombok.RequiredArgsConstructor;
import org.springframework.scheduling.annotation.Async;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.*;
import java.util.concurrent.CompletableFuture;

@Service
@RequiredArgsConstructor
@EnableAsync
public class CloudinaryService {

    private final Cloudinary cloudinary;

    @Async
    public void deleteExtraImage(List<MultipartFile> extraImage, Product product)
            throws FuncErrorException {
        if (extraImage == null || extraImage.isEmpty()) {
            CompletableFuture.completedFuture(null);
            return;
        }
        Set<ProductImage> extraImages = product.getExtraImages();
        for (ProductImage image : extraImages) {
            deleteImage(image.getImageId());
        }
        CompletableFuture.completedFuture(null);
    }

    @Async
    public void deleteMainImage(MultipartFile mainFile, Product product)
            throws FuncErrorException {
        if (mainFile != null && !mainFile.isEmpty()) {
            deleteImage(product.getImageId());
        }
        CompletableFuture.completedFuture(null);
    }


    public void setExtraImage(List<MultipartFile> extraFile,Product productInDB, RequestProduct productDTO )
            throws FuncErrorException {

        Set<CloudinaryDto> images = new HashSet<>();
        if (extraFile == null && extraFile.isEmpty() ) {
            if (productInDB != null ) {
                productInDB.getExtraImages().forEach(image ->
                    images.add(new CloudinaryDto(image.getImageId(),image.getUrl()))
                );
                productDTO.setExtraImages(images);
            } else {
                productDTO.setExtraImages(images);
            }
        } else {
            for (MultipartFile file : extraFile) {
                String fileName = FileUploadUtil.getFileName(file.getOriginalFilename());

                FileUploadUtil.assertAllowed(file, FileUploadUtil.IMAGE_PATTERN);

                CloudinaryDto result = uploadImage(file, fileName);
                images.add(result);
            }
            productDTO.setExtraImages(images);
        }
    }
//
//

    @Transactional(rollbackFor = {FuncErrorException.class})
    public void setMainImage(MultipartFile mainFile,Product product, RequestProduct productDTO)
            throws FuncErrorException {

        if (mainFile == null || mainFile.isEmpty() && product != null) {
            productDTO.setImageId(product.getImageId());
            productDTO.setUrl(product.getUrl());

        } else {
            FileUploadUtil.assertAllowed(mainFile, FileUploadUtil.IMAGE_PATTERN);

            String fileName = FileUploadUtil.getFileName(mainFile.getOriginalFilename());

            CloudinaryDto result = uploadImage(mainFile, fileName);

            productDTO.setImageId(result.getPublicId());
            productDTO.setUrl(result.getUrl());
        }

    }

    public CloudinaryDto uploadImage(MultipartFile file, String fileName) throws FuncErrorException {
        try {
            Map result = cloudinary.uploader()
                    .upload(file.getBytes(), Map.of("public_id", "microservice/product/" + fileName));
            String publicId = (String) result.get("public_id");
            String url = (String) result.get("secure_url");
            return new CloudinaryDto(publicId,url);
        } catch (Exception ex) {
            throw new FuncErrorException("File upload fail!");
        }
    }

    private void deleteImage(String publicId) {
        CompletableFuture.runAsync(() -> {
            try {
                cloudinary.uploader().destroy(publicId, null);
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        });

    }
}
