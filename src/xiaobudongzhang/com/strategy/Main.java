package xiaobudongzhang.com.strategy;

public class Main {
    public static void main(String[] args){
        //选择使用的策略
        Strategy s = new ConcreteStrategyA();
        Context context = new Context(s);
        context.contextInterface();
    }
}
