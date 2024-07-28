package com.phuc.categoryservice.dtos;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

import java.util.List;

@Setter
@Getter
@Builder
public class PaginationDto {
    @JsonProperty("total_items")
    private Integer totalItems;

    @JsonProperty("total_pages")
    private Integer totalPages;

    @JsonProperty("current_page")
    private Integer currentPage;

    @JsonProperty("start_count")
    private Integer startCount;

    @JsonProperty("end_count")
    private int endCount;

    @JsonProperty("entities")
    private List<?> entities;
}
