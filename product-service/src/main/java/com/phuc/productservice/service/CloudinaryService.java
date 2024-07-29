package com.phuc.productservice.service;

import com.cloudinary.Cloudinary;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.exceptions.FuncErrorException;
import com.phuc.productservice.request.RequestCreateProduct;
import com.phuc.productservice.util.FileUploadUtil;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Map;

@Service
@RequiredArgsConstructor
public class CloudinaryService {

    private final Cloudinary cloudinary;

    @Transactional(rollbackFor = {FuncErrorException.class})
    public void setExtraImage(MultipartFile[] extraFile, RequestCreateProduct requestCreateProduct) throws FuncErrorException {
        ArrayList<CloudinaryDto> images = new ArrayList<>();

        if (extraFile != null && extraFile.length > 0 ) {
            for (MultipartFile file : extraFile) {
                FileUploadUtil.assertAllowed(file, FileUploadUtil.IMAGE_PATTERN);

                String fileName = FileUploadUtil.getFileName(file.getOriginalFilename());

                CloudinaryDto result = upload(file, fileName);
                images.add(CloudinaryDto
                        .builder()
                        .url(result.getUrl())
                        .publicId(result.getPublicId()).build());
            }
        }
        requestCreateProduct.setExtraImages(images);
    }

    @Transactional(rollbackFor = {FuncErrorException.class})
    public void setMainImage(MultipartFile mainFile,RequestCreateProduct reqCreateProduct)
            throws FuncErrorException {
        if (mainFile != null && !mainFile.isEmpty()) {
            FileUploadUtil.assertAllowed(mainFile, FileUploadUtil.IMAGE_PATTERN);

            String fileName = FileUploadUtil.getFileName(mainFile.getOriginalFilename());

            CloudinaryDto result = upload(mainFile, fileName);

            reqCreateProduct.setImageId(result.getPublicId());
            reqCreateProduct.setUrl(result.getUrl());
        } else {
            reqCreateProduct.setImageId("");
            reqCreateProduct.setUrl("");
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
}
