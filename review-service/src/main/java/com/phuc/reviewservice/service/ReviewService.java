package com.phuc.reviewservice.service;

import com.phuc.reviewservice.constants.Constants;
import com.phuc.reviewservice.events.message.AvgRatingEvent;
import com.phuc.reviewservice.exeptions.DataErrorException;
import com.phuc.reviewservice.exeptions.ParamValidateException;
import com.phuc.reviewservice.models.Review;
import com.phuc.reviewservice.repository.ReviewRepository;
import com.phuc.reviewservice.request.RequestReview;
import com.phuc.reviewservice.util.Utility;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;


@Service
@RequiredArgsConstructor
public class ReviewService implements IReviewService{

    private final ReviewRepository reviewRepository;
    private final KafkaTemplate<String, Object> kafkaTemplate;

    public Review deleteReview(String reviewId) throws DataErrorException {
        Review review = reviewRepository.findById(reviewId).orElseThrow(() ->
                new DataErrorException(Constants.DB_NOT_FOUND)
        );

        reviewRepository.delete(review);

        return review;
    }
    public void postReview(String customerId,String proId, RequestReview reqReview) {

        Review review = Review.builder()
                .customerId(customerId)
                .productId(proId)
                .headline(reqReview.getHeadline())
                .comment(reqReview.getComment())
                .rating(reqReview.getRating()).build();
        reviewRepository.save(review);

        Float avgRating = reviewRepository.findAvgRating(proId);
        AvgRatingEvent avg = AvgRatingEvent.builder()
                .averageRating(avgRating)
                .productId(proId).build();

        kafkaTemplate.send(Constants.TOPIC_AVG, avg);
    }

    public Page<Review> getAllReviews(Integer page,Integer limit,String sort, String keyword) throws ParamValidateException {
        Utility.checkSortIsAscOrDesc(sort);

        Sort sortDir = Sort.by("createdAt");
        sortDir = sort.equals("asc") ? sortDir.ascending() : sortDir.descending();

        Pageable pageable = PageRequest.of(page - 1,limit,sortDir);

        if (!keyword.isEmpty()) {
            return reviewRepository.search(keyword, pageable);
        } else  {
            return reviewRepository.findAll(pageable);
        }
    }

    public Page<Review> getAllReviewsByProduct(String proId, Integer page,Integer limit, Integer rating) throws ParamValidateException {
        if (rating < 1 || rating > 5) {
            throw new ParamValidateException("rating < 1 or > 5!");
        }

        Sort sortDir = Sort.by("rating").descending();
        Pageable pageable = PageRequest.of(page - 1,limit,sortDir);

        return reviewRepository.findByProduct(proId,pageable);
    }
}
