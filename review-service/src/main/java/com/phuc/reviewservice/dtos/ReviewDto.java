package com.phuc.reviewservice.dtos;


import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
public class ReviewDto {
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
}

