#
# Thanks for your contribution !!!
#

name: 'Awesome Pull Request :rocket:'  
description: Make a change to have go-echarts better.  
body:
  - type: textarea  
    attributes:  
    label: Description  
    description: A description of your improvement changes.  

  - type: dropdown
    attributes:
      label: Changes Type  
      multiple: true  
      options:  
        - 'Brand New Charts'  
        - 'Enrich Options'  
        - 'Bug Fix'  
        - 'New Feature'  
        - 'Other'  
    validations:   
      required: true  

  - type: checkboxes  
    attributes:  
      label: Are you willing to submit a PR on [Examples](https://github.com/go-echarts/examples)?  
    description: >  
        This is absolutely not required, but we are happy to see that you could share or update the related
        charts examples to more users.  
    options:  
      - label: Yes, I am willing to submit a PR on [Examples](https://github.com/go-echarts/examples))!

  - type: markdown  
    attributes:  
      value: "Thanks for your contributions !!!"  