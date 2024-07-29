package com.phuc.productservice.request;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonManagedReference;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.phuc.productservice.dtos.CloudinaryDto;
import com.phuc.productservice.models.ProductImage;
import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.OneToMany;
import jakarta.validation.constraints.*;
import lombok.Getter;
import lombok.Setter;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

@Getter
@Setter
public class RequestCreateProduct {
    @JsonProperty("name")
    @NotBlank(message = "name cannot empty or null")
    private String name;

    @JsonProperty("description")
    @NotBlank(message = "description cannot empty or null")
    @Size(max = 10000, message = "description cannot be longer than 10000 characters")
    private String description;

    @JsonProperty("status")
    @NotBlank(message = "status must be true or false")
    private Boolean status;

    @JsonProperty("cost")
    @NotNull(message = "cost cannot be null")
    @Min(value = 0, message = "cost cannot be less than 0")
    private Float cost;

    @JsonProperty("price")
    @NotNull(message = "price cannot be null")
    @Min(value = 0, message = "price cannot be less than 0")
    private Float price;

    @JsonProperty("sale")
    @NotNull(message = "sale cannot be null")
    @Min(value = 0, message = "sale cannot be less than 0")
    @Max(value = 100, message = "sale cannot greater than 0")
    private Float sale;

    @JsonProperty("category_id")
    @NotNull(message = "category_id cannot null")
    private String categoryId;

    @JsonIgnore
    private String imageId;
    @JsonIgnore
    private String url;
    @JsonIgnore
    private List<CloudinaryDto> extraImages = new ArrayList<>();
}
