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
    private int xp;
    private byte[] photo;

    public UsersEntity() {
    }

    public UsersEntity(Integer id, String name, int xp, byte[] photo) {
        this.id = id;
        this.name = name;
        this.xp = xp;
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

    public int getXp() {
        return xp;
    }

    public void setXp(int xp) {
        this.xp = xp;
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
                ", xp=" + xp +
                ", photo=" + Arrays.toString(photo) +
                '}';
    }
}
