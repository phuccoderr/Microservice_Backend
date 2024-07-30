package com.phuc.productservice.service;

import com.cloudinary.Cloudinary;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.FuncErrorException;
import com.phuc.productservice.models.Product;
import com.phuc.productservice.models.ProductImage;
import com.phuc.productservice.request.RequestCreateProduct;
import com.phuc.productservice.request.RequestUpdateProduct;
import com.phuc.productservice.util.FileUploadUtil;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

@Service
@RequiredArgsConstructor
public class CloudinaryService {

    private final Cloudinary cloudinary;

    @Transactional(rollbackFor = {FuncErrorException.class})
    public void deleteExtraImage(MultipartFile[] extraImage, Product product)
            throws FuncErrorException {
        if (extraImage == null || extraImage.length == 0) return;
        Set<ProductImage> extraImages = product.getExtraImages();
        for (ProductImage image : extraImages) {
            deleteImage(image.getImageId());
        }
    }

    @Transactional(rollbackFor = {FuncErrorException.class})
    public void deleteMainImage(MultipartFile mainFile, Product product)
            throws FuncErrorException {
        if (mainFile != null && !mainFile.isEmpty()) {
            deleteImage(product.getImageId());
        }

    }

    @Transactional(rollbackFor = {FuncErrorException.class})
    public void setExtraImage(MultipartFile[] extraFile, Product product)
            throws FuncErrorException {

        if (extraFile != null && extraFile.length > 0 ) {
            for (MultipartFile file : extraFile) {
                String fileName = FileUploadUtil.getFileName(file.getOriginalFilename());

                FileUploadUtil.assertAllowed(file, FileUploadUtil.IMAGE_PATTERN);

                CloudinaryDto result = upload(file, fileName);
                product.addImage(result);
            }
        }
    }

    @Transactional(rollbackFor = {FuncErrorException.class})
    public void setMainImage(MultipartFile mainFile,Product product)
            throws FuncErrorException {
        if (mainFile != null && !mainFile.isEmpty()) {
            FileUploadUtil.assertAllowed(mainFile, FileUploadUtil.IMAGE_PATTERN);

            String fileName = FileUploadUtil.getFileName(mainFile.getOriginalFilename());

            CloudinaryDto result = upload(mainFile, fileName);

            product.setImageId(result.getPublicId());
            product.setUrl(result.getUrl());
        }
        if (product.getImageId() == null || product.getImageId().isEmpty()) {
            product.setImageId("");
            product.setUrl("default image!");
        }
    }

    private CloudinaryDto upload(MultipartFile file, String fileName) throws FuncErrorException {
        try {
            Map result = cloudinary.uploader()
                    .upload(file.getBytes(), Map.of("public_id", "microservice/product/" + fileName));
            String url = (String) result.get("secure_url");
            String publicId = (String) result.get("public_id");
            return CloudinaryDto.builder().url(url).publicId(publicId).build();
        } catch (Exception ex) {
            throw new FuncErrorException("Failed to upload file");
        }
    }

    private void deleteImage(String publicId) throws FuncErrorException {
        try {
            cloudinary.uploader().destroy(publicId, null);
        } catch (IOException e) {
            throw new FuncErrorException("Image delete fail");
        }
    }

    private String getPublicIdFromUrl(String imageUrl) {
        // https://res.cloudinary.com/dp4tp9gwa/image/upload/v1713253619/b8db51d6-d20b-420d-860e-3228352026f4.jpg
        String[] parts = imageUrl.split("/");
        String publicIdWithExtension = parts[parts.length - 1]; // b8db51d6-d20b-420d-860e-3228352026f4.jpg
        String[] publicIdParts = publicIdWithExtension.split("\\."); // b8db51d6-d20b-420d-860e-3228352026f4
        return publicIdParts[0]; // Public_id là phần trước dấu chấm trong publicIdWithExtension
    }
}
