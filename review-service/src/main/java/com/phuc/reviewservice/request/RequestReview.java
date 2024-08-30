package com.phuc.reviewservice.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Positive;
import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
public class RequestReview {
    @JsonProperty("headline")
    @NotBlank(message = "headline cannot null or empty")
    private String headline;

    @JsonProperty("comment")
    @NotBlank(message = "comment cannot null or empty")
    private String comment;

    @JsonProperty("rating")
    @Min(value = 1, message = "rating > 1 or <5")
    @Max(value = 5, message = "rating > 1 or <5")
    private int rating;
}
