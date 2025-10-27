package dustbox.dustbox.model

package com.example.model;

@Entity
@Table(name = "users")
public class User 
{
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(nullable = false)
    private String name;
    private String email;
    
    public User(String name, String email) 
    {
        this.name = name;
        this.email = email;
    }
}
