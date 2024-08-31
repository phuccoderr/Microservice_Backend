package com.phuc.reviewservice.service;

import com.phuc.reviewservice.exeptions.DataErrorException;
import com.phuc.reviewservice.exeptions.ParamValidateException;
import com.phuc.reviewservice.models.Review;
import com.phuc.reviewservice.request.RequestReview;
import org.springframework.data.domain.Page;

public interface IReviewService {
    void postReview(String customerId,String proId, RequestReview reqReview);
    Review deleteReview(String reviewId) throws DataErrorException;
    Page<Review> getAllReviews(Integer page, Integer limit, String sort, String keyword) throws ParamValidateException;
    Page<Review> getAllReviewsByProduct(String proId, Integer page,Integer limit, Integer rating) throws ParamValidateException;
}
