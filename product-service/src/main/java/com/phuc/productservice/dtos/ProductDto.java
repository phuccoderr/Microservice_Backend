package com.phuc.productservice.dtos;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.phuc.productservice.models.ProductImage;
import lombok.*;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

@Setter
@Getter
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class ProductDto {
    @JsonProperty("id")
    private String id;

    @JsonProperty("created_at")
    private LocalDateTime createdAt;

    @JsonProperty("updated_at")
    private LocalDateTime updatedAt;

    @JsonProperty("name")
    private String name;

    @JsonProperty("alias")
    private String alias;

    @JsonProperty("description")
    private String description;

    @JsonProperty("status")
    private Boolean status;

    @JsonProperty("cost")
    private Float cost;

    @JsonProperty("price")
    private Float price;

    @JsonProperty("sale")
    private Float sale;

    @JsonProperty("url")
    private String url;

    @JsonProperty("extra_images")
    private List<String> extraImages = new ArrayList<>();

    @JsonProperty("category_id")
    private String categoryId;
}
