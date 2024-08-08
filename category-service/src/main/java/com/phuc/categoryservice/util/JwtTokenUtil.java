package com.phuc.categoryservice.util;

import com.phuc.categoryservice.constants.Constants;
import com.phuc.categoryservice.exceptions.InvalidTokenException;
import com.phuc.categoryservice.exceptions.TokenExpiredException;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.ExpiredJwtException;
import io.jsonwebtoken.Jwts;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class JwtTokenUtil {

    @Value("${JWT_SECRET}")
    private String secretKey;

    public Claims extractToken(String token) {
        try {
            return Jwts.parser()
                    .setSigningKey(secretKey.getBytes())
                    .build()
                    .parseSignedClaims(token)
                    .getPayload();
        } catch (ExpiredJwtException e) {
            throw new TokenExpiredException(Constants.TOKEN_EXPIRED);
        } catch (Exception e) {
            throw new InvalidTokenException(Constants.TOKEN_INVALID);
        }
    }

}
