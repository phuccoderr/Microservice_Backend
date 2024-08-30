package com.phuc.reviewservice.events.message;


import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

import java.io.Serializable;

@Setter
@Getter
@Builder
public class AvgRatingEvent implements Serializable {
    private static final long serialVersionUID = 1L;
    @JsonProperty("product_id")
    private String productId;

    @JsonProperty("average_rating")
    private Float averageRating;
}

