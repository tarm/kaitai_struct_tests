using NUnit.Framework;

namespace Kaitai
{
    [TestFixture]
    public class SpecExpr0 : CommonSpec
    {
        [Test]
        public void TestExpr0()
        {
            Expr0 r = Expr0.FromFile(SourceFile("str_encodings.bin"));

            Assert.AreEqual(r.MustBeF7, 0xf7);
            Assert.AreEqual(r.MustBeAbc123, "abc123");
        }
    }
}
