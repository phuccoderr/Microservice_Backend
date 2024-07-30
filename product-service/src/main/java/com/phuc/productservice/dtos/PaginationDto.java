package com.phuc.productservice.dtos;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.phuc.productservice.models.Product;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.data.domain.Page;

import java.util.List;

@Setter
@Getter
@NoArgsConstructor
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

    public PaginationDto(Page<Product> pages, List<ProductDto> listDtos) {
        int pageSize = pages.getSize();
        int oldCurrentPage = pages.getNumber();

        this.currentPage = oldCurrentPage + 1;
        this.totalPages = pages.getNumber();
        this.totalItems = (int) pages.getTotalElements();
        this.startCount = (oldCurrentPage * pageSize) + 1;
        this.endCount = Math.min((oldCurrentPage * pageSize) + pageSize, totalItems);
        this.entities = listDtos;
    }
}
