package xiaobudongzhang.com.strategy;

public class Context {
    //持有一个具体策略的对象
    private Strategy strategy;

    public Context(Strategy strategy){
        this.strategy = strategy;
    }

    public void contextInterface(){
        strategy.algorithmInterface();
    }
}
