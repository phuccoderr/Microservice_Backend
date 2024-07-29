package com.phuc.categoryservice.util;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class JwtTokenUtil {

    @Value("${JWT_SECRET}")
    private String secretKey;

    public Claims extractToken(String token) {
        return Jwts.parser()
                .setSigningKey(secretKey.getBytes())
                .build()
                .parseSignedClaims(token)
                .getPayload();
    }

}
