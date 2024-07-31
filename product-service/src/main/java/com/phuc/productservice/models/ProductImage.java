package com.phuc.productservice.models;

import com.fasterxml.jackson.annotation.JsonBackReference;
import com.phuc.productservice.dtos.CloudinaryDto;
import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.Setter;

@Setter
@Getter
@AllArgsConstructor
@RequiredArgsConstructor
@Entity
@Table(name = "product_images")
public class ProductImage extends BaseEntity {
    @Column(name = "image_id")
    private String imageId;
    private String url;

    @ManyToOne
    @JoinColumn(name = "product_id")
    @JsonBackReference
    private Product product;

    public ProductImage(CloudinaryDto cloudinaryDto, Product product) {
        this.imageId = cloudinaryDto.getPublicId();
        this.url = cloudinaryDto.getUrl();
        this.product = product;
    }

    public ProductImage(CloudinaryDto cloudinaryDto) {
        this.imageId = cloudinaryDto.getPublicId();
        this.url = cloudinaryDto.getUrl();
    }
}
