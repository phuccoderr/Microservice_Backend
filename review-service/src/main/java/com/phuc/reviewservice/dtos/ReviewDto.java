package com.phuc.reviewservice.dtos;


import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Getter;
import lombok.Setter;

import java.time.LocalDateTime;

@Setter
@Getter
public class ReviewDto {

    @JsonProperty("id")
    private String id;

    @JsonProperty("name")
    private String name;

    @JsonProperty("headline")
    private String headline;

    @JsonProperty("comment")
    private String comment;

    @JsonProperty("rating")
    private int rating;

    @JsonProperty("product_id")
    private String productId;

    @JsonProperty("customer_id")
    private String customerId;

    @JsonProperty("created_at")
    private LocalDateTime createdAt;

    @JsonProperty("updated_at")
    private LocalDateTime updatedAt;
}

