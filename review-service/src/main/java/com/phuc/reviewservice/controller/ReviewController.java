package com.phuc.reviewservice.controller;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.phuc.reviewservice.constants.Constants;
import com.phuc.reviewservice.dtos.PaginationDto;
import com.phuc.reviewservice.dtos.ReviewDto;
import com.phuc.reviewservice.exeptions.DataErrorException;
import com.phuc.reviewservice.exeptions.ParamValidateException;
import com.phuc.reviewservice.models.Review;
import com.phuc.reviewservice.request.RequestReview;
import com.phuc.reviewservice.response.ResponseObject;
import com.phuc.reviewservice.service.ProductService;
import com.phuc.reviewservice.service.ReviewRedisService;
import com.phuc.reviewservice.service.ReviewService;
import com.phuc.reviewservice.util.JwtTokenUtil;
import com.phuc.reviewservice.util.Utility;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

import java.util.List;


@RestController
@RequestMapping(Constants.API_REVIEWS)
@RequiredArgsConstructor
public class ReviewController {

    private final ProductService productService;
    private final ReviewService reviewService;
    private final ReviewRedisService reviewRedisService;
    private final JwtTokenUtil jwtTokenUtil;

    @GetMapping("")
    public ResponseEntity<ResponseObject> listByPage(
            @RequestParam(value = "page", defaultValue = "1") Integer page,
            @RequestParam(value = "limit", defaultValue = "10") Integer limit,
            @RequestParam(value = "sort", defaultValue = "asc") String sort,
            @RequestParam(value = "keyword", defaultValue = "") String keyword
    ) throws ParamValidateException, JsonProcessingException {

        if (keyword.isEmpty()) {
            PaginationDto paginationDto = reviewRedisService.getAllReviews(page, limit, sort);
            if (paginationDto != null ) {
                return new ResponseEntity<>(ResponseObject.builder()
                        .status(HttpStatus.OK.value())
                        .message(Constants.GET_ALL_SUCCESS)
                        .data(paginationDto).build(), HttpStatus.OK);
            }
        }
        Page<Review> pages = reviewService.getAllReviews(page, limit, sort, keyword);

        List<ReviewDto> listDtos = Utility.toListDtos(pages.getContent());

        PaginationDto paginationDto = new PaginationDto(pages,listDtos);

        if (keyword.isEmpty()) {
            reviewRedisService.saveAllReviews(paginationDto,page, limit, sort);
        }

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.GET_ALL_SUCCESS)
                .data(paginationDto).build(), HttpStatus.OK);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<ResponseObject> deleteReview(
            @PathVariable("id") String reviewId
    ) throws DataErrorException {
        Review review = reviewService.deleteReview(reviewId);

        reviewRedisService.clearAllReviews();
        reviewRedisService.clearAllReviewsByProduct(review.getProductId());

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.DELETE_SUCCESS)
                .data(null).build(), HttpStatus.OK);
    }

    @GetMapping("/ratings/{proId}")
    public ResponseEntity<ResponseObject> listRating(
        @PathVariable("proId") String proId,
        @RequestParam(value = "page", defaultValue = "1") Integer page,
        @RequestParam(value = "limit", defaultValue = "10") Integer limit,
        @RequestParam(value = "rating", defaultValue = "5") Integer rating,
        HttpServletRequest request
    ) throws ParamValidateException, JsonProcessingException {

        if(rating == 5) {
            PaginationDto paginationDto = reviewRedisService.getAllReviewsByProduct(proId, page, limit);
            if (paginationDto != null ) {
                return new ResponseEntity<>(ResponseObject.builder()
                        .status(HttpStatus.OK.value())
                        .message(Constants.GET_ALL_SUCCESS)
                        .data(paginationDto).build(), HttpStatus.OK);
            }
        }

        productService.findById(proId,request);

        Page<Review> pages = reviewService.getAllReviewsByProduct(proId, page, limit, rating);
        List<ReviewDto> listDtos = Utility.toListDtos(pages.getContent());

        PaginationDto paginationDto = new PaginationDto(pages,listDtos);
        reviewRedisService.saveAllReviewsByProduct(paginationDto,proId,page,limit);

        return new ResponseEntity<>(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message(Constants.GET_ALL_SUCCESS)
                .data(paginationDto).build(), HttpStatus.OK);
    }

    @PostMapping("/post_review/{id}")
    public ResponseEntity<ResponseObject> postReview(@PathVariable("id") String proId,
                                                     @RequestBody @Valid RequestReview review ) {

        String customerId = jwtTokenUtil.getCustomerId(SecurityContextHolder.getContext().getAuthentication());
        String customerName = jwtTokenUtil.getNameCustomer(SecurityContextHolder.getContext().getAuthentication());

        reviewService.postReview(customerId, customerName,proId,review);
        reviewRedisService.clearAllReviews();
        reviewRedisService.clearAllReviewsByProduct(proId);

        return ResponseEntity.ok(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("OK")
                .data(proId).build());
    }

    @GetMapping("/can_review")
    public ResponseEntity<ResponseObject> canReview(
            @RequestParam(value = "product_id") String proId,
            @RequestParam(value = "customer_id") String customerId) throws DataErrorException {
        Review review = reviewService.canCustomerReviewProduct(customerId, proId);
        ReviewDto dto = Utility.toDto(review);

        return ResponseEntity.ok(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("OK")
                .data(dto).build());

    }

    @GetMapping("/{id}")
    public ResponseEntity<ResponseObject> getOne(
            @PathVariable("id") String id
    ) throws DataErrorException {
        Review review = reviewService.getById(id);
        ReviewDto dto = Utility.toDto(review);

        return ResponseEntity.ok(ResponseObject.builder()
                .status(HttpStatus.OK.value())
                .message("OK")
                .data(dto).build());
    }
}
