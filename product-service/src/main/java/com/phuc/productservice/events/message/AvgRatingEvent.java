package com.phuc.productservice.events.message;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
public class AvgRatingEvent {
    @JsonProperty("product_id")
    private String productId;

    @JsonProperty("average_rating")
    private Float averageRating;
}
