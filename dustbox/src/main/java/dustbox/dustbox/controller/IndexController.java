package dustbox.dustbox;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;

@RestController
public class IndexController 
{
  @GetMapping("/")
  public String index() 
  {
    return "jsjsj";
  }
}
