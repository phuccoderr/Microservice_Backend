package com.phuc.productservice.models;


import com.fasterxml.jackson.annotation.JsonManagedReference;
import com.phuc.productservice.dtos.CloudinaryDto;
import jakarta.persistence.*;
import lombok.*;

import java.util.HashSet;
import java.util.Set;

@Setter
@Getter
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "products")
public class Product extends BaseEntity{
    @Column(unique = true)
    private String name;

    @Column(unique = true)
    private String alias;

    private String description;
    private Boolean status;
    private Integer stock;
    private Float cost;
    private Float price;
    private Float sale;

    @Column(name = "image_id")
    private String imageId;
    private String url;

    @OneToMany(mappedBy = "product", cascade = CascadeType.ALL, orphanRemoval = true)
    @JsonManagedReference
    private Set<ProductImage> extraImages = new HashSet<>();

    @Column(name = "category_id")
    private String categoryId;

    @Column(name = "average_rating")
    private Float averageRating;

    @Column(name = "review_count")
    private Integer reviewCount;

    public void addImage(CloudinaryDto cloudinaryDto) {
        if (extraImages == null) {
            extraImages = new HashSet<>();
        }
        this.extraImages.add(new ProductImage(cloudinaryDto,this));
    }

}
