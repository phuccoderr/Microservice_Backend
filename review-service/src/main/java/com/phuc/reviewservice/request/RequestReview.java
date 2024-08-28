package com.phuc.reviewservice.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
public class RequestReview {
    @JsonProperty("headline")
    private String headline;

    @JsonProperty("comment")
    private String comment;

    @JsonProperty("rating")
    private int rating;
}
