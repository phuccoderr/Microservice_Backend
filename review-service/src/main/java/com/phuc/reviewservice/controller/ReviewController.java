package com.phuc.reviewservice.controller;

import com.phuc.reviewservice.constants.Constants;
import com.phuc.reviewservice.request.RequestReview;
import com.phuc.reviewservice.response.ResponseObject;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(Constants.API_CATEGORIES)
public class ReviewController {

    @PostMapping("/post_review/{id}")
    public ResponseEntity<ResponseObject> postReview(@PathVariable("id") String proId,
                                                     @RequestBody RequestReview review) {
        return null;
    }
}
