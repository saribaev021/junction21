package com.app.UserInfo.entity;

import javax.persistence.*;
import java.util.Arrays;

@Entity
@Table(name = "users")
public class UsersEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    private String name;
    private double rating;
    private byte[] photo;

    public UsersEntity() {
    }

    public UsersEntity(Integer id, String name, double rating, byte[] photo) {
        this.id = id;
        this.name = name;
        this.rating = rating;
        this.photo = photo;
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public double getRating() {
        return rating;
    }

    public void setRating(double rating) {
        this.rating = rating;
    }

    public byte[] getPhoto() {
        return photo;
    }

    public void setPhoto(byte[] photo) {
        this.photo = photo;
    }

    @Override
    public String toString() {
        return "UserEntity{" +
                "id=" + id +
                ", name='" + name + '\'' +
                ", rating=" + rating +
                ", photo=" + Arrays.toString(photo) +
                '}';
    }
}
