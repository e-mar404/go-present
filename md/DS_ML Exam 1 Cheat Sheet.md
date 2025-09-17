#### Prediction & Inference:
*Prediction* deals with predicting an outcome without worrying about the relationship of variables that would impact it.
*Inference* deals more with the factors that affect something, trying to reach an explanation in a response variable.

#### Quantitative or Qualitative: 
*Quantitative* : variable that can be explained in a numerical sense.
*Qualitative* : variable that is part of a category, class or a factor. (Quality) 
To use *qualitative* variables we turn them into dummy variables so we can perform mathematical processed with them.

###### Rcode:
To transform a variable into a dummy var we use the following code:  `x <- ifelse(x == "yes", 1, 0)` this will make x = 1 if x == "yes" and 0 otherwise.
If you are wanting just to make a qualitative variable then: `x <- as.factor(as.character(x))`
If wanting to convert a qualitative var to quantitative: `x <- as.numeric(x)`

#### Regression and Classification:
We use *regression* when the response variable is `quantitative` and we use *classification* when the respoinse variable is `qualitative`.
- Regression: (Multiple)/Linear regression
- Classification: Logistic Regression

#### Linear Regression:
###### Formula Review:
$$\text{Var} = \sigma^2/n = \text{MSE, if Bias is 0} $$
$$\text{Bias}(\hat{\theta}) = E(\hat{\theta}) - \theta$$
$$\text{MSE}(\hat{\theta} - \theta)^2 = \text{Var}(\hat{\theta}) +[\text{Bias}(\hat{\theta})]^2 = \text{SSE}/(n-p-1)$$
$$\text{residual} = y_i - \hat{y_i}$$
$$\text{AIC} = \text{residual deviance} + 2\times(p+1) \text{ smaller is better}$$
$$\text{accuracy} = (\text{True Positive} + \text{True Negative})/Total$$
$$C_p = \text{SSE}/\text{MSE } + 2(p+1)-n \text{ smaller is better}$$
$$R^2 = 1 - (SSE(\text{residual deviance})/TSS(\text{null deviance})) \text{ higher is better}$$
$$\text{Adjusted-}R^2 = 1 - (SSE/(n-p-1)/(SST/(n-1))$$
$$\text{sensitivity} = \text{True Positives}/\text{Positives}$$
$$\text{specificity} = \text{True Negative}/\text{Negative}$$
$$\text{SSE = Sum Sq for res \& SST - Sum Sq total \& SST = Sum Sq all}$$
#### Linear Regression:
`Pr(>|t|)` is the p-value for the following Test Hypothesis:
	$H_0: \beta_i = 0$ and $H_a: \beta_i \neq 0$ given that the other variables are in the model.

$R^2$  is the % of variation that can be explained by the model.
`F-statistic` is the test state for the following Test Hypothesis:
$H_0: \beta_1 = \beta_2 = ... = \beta_p = 0$ and $H_a: \text{At least one }\beta_j = 0$
Meaning that there is atleast one predictor that is significant in the model.

To find prediction/confidence interval use the function: `predict(obj.lm, newdata=data.frame(x=#), interval="c"/"p", level=0.95`
*Prediction Interval* : predicts an individual value, it is a wider range therfore less certain
*Confidence Interval*: predicts the mean value for any given obs, smaller range and more certain since it looks at other observations
*Linear Discriminant Analysis*: has to do with classes instead of numbers

#### LINE:
To use linear regression we need to make sure of the following: *Linear, INdependent, Normal Distribution, Equal Variance*

#### Step:
*Additive Regression*: look at all possible models to determine which is best. Ther are $2^p$ models trhat contains subsets of p variables(predictors)
Use code `step(obj.lm, direction=c("both", "forward", "backward"))`

###### R code/Ploting:
Use `pairs()` to make a scatter plot with variables passes to the function.
Accessing col and row `x[col,row]`. i.e: `x[,3:6]` are rows 3 to 6 from data set x to get rid of col or row just make negative.
If wanting to look at linear model to check *LINE* then use: `plot(obj.lm)` and `par(mfrow = c(2,2))`

#### Polynomial Regression:
If when using `pairs()` you notice a non-linear relationship use *Polynomial Regression*. This helps make data linear and normal.
To transform a variable by power of x or $\sqrt{x}$  we would do so like: `lm(y ~ poly(a,2) + sqrt(b))` 

#### Logistic Regression:
We use *logistic regression* when the response variable is *qualitative*. We also need to do transformations:
- $p(x) = P(Y = 1| x)$ this will model the probability of $Y = 1$ 
- $p(x) \in [0, 1]$
- formula $ln(p(x)/(1/p(x)))$ is known as `logit` 
To call *logistic regression* use: `glm(y ~ x, family="binomial", data=)`

#### Confusion Matrix:
|         | $\hat{Y} = 0$             | $\hat{Y} = 1$             |
| ------- | ------------------------- | ------------------------- |
| $Y = 0$ | Correct/True Negatives    | Incorrect/False Positives |
| $Y = 1$ | Incorrect/False Negatives | Correct/True Positives    |